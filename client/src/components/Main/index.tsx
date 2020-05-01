import React, { useContext } from 'react';
import './main.scss';
import Login from '../Login';
import Chat from '../Chat';
import ChatContext, { ChatContextInterface } from '../../contexts/chat';

function Main(): JSX.Element {
  const { author } = useContext<ChatContextInterface>(ChatContext);

  return (
    <div>
      <header id="header">
        <h1>Gabble</h1>
        <span>💬 Tiny chat implementation made with Go and ReactJS</span>
      </header>
      <main>
        {
          author ? <Chat /> : <Login />
        }
      </main>
    </div>
  );
}

export default Main;
