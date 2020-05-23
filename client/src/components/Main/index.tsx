import React, { useContext } from 'react';
import './main.scss';
import Login from '../Login';
import Chat from '../Chat';
import ChatContext, { IChatContext } from '../../contexts/chat';

function Main(): JSX.Element {
  const { isConnected } = useContext<IChatContext>(ChatContext);

  return (
    <div>
      <header id="header">
        <h1>Gabble</h1>
        <span>💬 Tiny chat implementation made with Go and ReactJS</span>
      </header>
      <main>
        {
          isConnected ? <Chat /> : <Login />
        }
      </main>
    </div>
  );
}

export default Main;
