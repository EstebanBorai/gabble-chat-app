import React from 'react';
import './chat.scss';
import ChatContext, { ChatContextInterface, Message } from '../../contexts/chat';
import Bubble from './Bubble';

const Chat = (): JSX.Element => {
  const [text, setText] = React.useState('');

  const { send, messages, author } = React.useContext<ChatContextInterface>(ChatContext);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setText(event.target.value);
  }

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();

    send(text);
  }

  return (
    <section className="application-section" id="chat-session">
      <ul className="chat">
        {
          messages.map((m: Message, index) => (
          <Bubble key={index} message={m} me={author} />
          ))
        }
      </ul>
      <div className="chat-input">
        <form action="" onSubmit={handleSubmit}>
          <input type="text" name="text" value={text} onChange={handleChange} />
          <button type="submit">Send</button>
        </form>
      </div>
    </section>
  );
}

export default Chat;
