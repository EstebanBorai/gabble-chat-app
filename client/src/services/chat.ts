import makeArrBuffStr from 'arrbuffstr';
import {
  BehaviorSubject, fromEvent, merge, Observable,
} from 'rxjs';
import { map, tap } from 'rxjs/operators';

export interface Message {
  body: any;
  author: string;
}

export type Stream = Observable<Message>;

export type ConnectionStatus = BehaviorSubject<boolean>;

export interface IChatService {
  stream: Stream;
  isConnected: ConnectionStatus;
  send: (message: string) => void;
  connect: (url: string) => Promise<void>;
  disconnect: () => void;
}

class ChatService implements IChatService {
  public stream: Stream;
  public isConnected: ConnectionStatus;
  private ws: WebSocket | null;
  private arrBuffStr: any;

  constructor() {
    this.ws = null;
    this.stream = new Observable<Message>();
    this.isConnected = new BehaviorSubject<boolean>(false);
    this.arrBuffStr = makeArrBuffStr();
  }

  public async connect(url: string): Promise<void> {
    return new Promise((resolve, reject): void => {
      try {
        this.ws = new WebSocket(url);
        this.isConnected.next(true);
        const closed$ = fromEvent(this.ws, 'close')
          .pipe(
            tap(() => {
              this.isConnected.next(false);
            }),
            map(() => ({
                body: 'Connection done',
                author: 'System'
              })));

        const message$ = fromEvent(this.ws, 'message')
          .pipe(
            map((event: MessageEvent): Message => ({
              body: event?.data,
              author: 'Another User',
            })));

        const send$ = fromEvent(this.ws, 'send')
          .pipe(
            map((event: CustomEvent): Message => ({
              body: event?.detail,
              author: 'Me'
            })));

        const open$ = fromEvent(this.ws, 'open')
        .pipe(
          tap(() => {
            this.isConnected.next(true);
          }),
          map(() => ({
              body: 'Connection established',
              author: 'System'
            })));

        const error = fromEvent(this.ws, 'error')
          .pipe(
            tap((e) => {
              console.error('WS ERROR', e);
            })
          )

        this.stream = merge(closed$, message$, send$, open$);

        return resolve();
      } catch (error) {
        console.error(error);
      }
    });
  }

  public disconnect(): void {
    this.ws.close();
  }

  public send(message: string): void {
    const event = new CustomEvent('send', {
      detail: message,
    });

    this.ws.send(this.arrBuffStr.toArrayBuffer(message));
    this.ws.dispatchEvent(event);
  }
}

export default ChatService;
