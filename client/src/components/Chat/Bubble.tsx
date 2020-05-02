import React from 'react';
import { Message, Author } from '../../services/ChatService';

interface BubbleProps {
  message: Message;
  me: Author;
}

const Bubble = ({ message, me }: BubbleProps): JSX.Element => {
  const isOwner = message.author && message.author.username === me.username;
  const isSystem = !message.author;

  if (isSystem) {
    return (
      <li className="bubble-container system">
        <article className="bubble">
          <p>
            {message.message}
          </p>
        </article>
      </li>
    );
  }

  return (
    <li className={`bubble-container${isOwner ? ' owner' : ''}`}>
      <article className="bubble">
        <small>{message.author.username}</small>
        <p>
          {message.message}
        </p>
      </article>
      <time>{message.issuedAt}</time>
    </li>
  );
}

export default Bubble;
