import { BehaviorSubject } from 'rxjs';
import StringBytes, { IStringBytes } from '../helpers/string-bytes';

export interface Author {
  username: string;
}

export interface Message {
  author: Author;
  message: string;
  issuedAt: Date;
}

export interface ChatServiceInterface {
  author: BehaviorSubject<Author | null>;
  messages: BehaviorSubject<Message[]>;
  send: (message: string) => void;
  join: (username: string) => void;
}

class ChatService implements ChatServiceInterface {
  private ws: WebSocket | null;
  private _messages: Message[];
  private _author: BehaviorSubject<Author | null>;
  private messageObservable: BehaviorSubject<Message[]>;

  constructor() {
    this.ws = null;
    this._author = new BehaviorSubject(null);
    this._messages = [];
    this.messageObservable = new BehaviorSubject([]);
  }

  public get author() {
    return this._author;
  }

  public get messages() {
    return this.messageObservable;
  }

  private buildMessage(message: string): ArrayBuffer {
    const raw = JSON.stringify({
      author: this.author.getValue().username,
      message: message,
      issuedAt: new Date().toISOString()
    });

    const stringBytes: IStringBytes = new StringBytes(raw);

    return stringBytes.toArrayBuffer();
  }

  public send = (message: string) => {
    if (!this._author.value) {
      throw new Error('An author is required in order to send a message');
    }

    (this.ws as WebSocket).send(this.buildMessage(message));
  }

  public join = (username: string) => {
    this._author.next({
      username
    });

    this.ws = new WebSocket(`ws://${process.env.WEB_SOCKET_HOST}:${process.env.WEB_SOCKET_PORT}`);

    this.ws.onmessage = (ev: MessageEvent): void => {
      this._messages.push(JSON.parse(ev.data));
      this.messageObservable.next(this._messages);
    }

    this.ws.onerror = (ev: Event): void => {
      console.log(ev);
    }
  }
}

export default ChatService;
