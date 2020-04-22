import React from 'react';
import './app.scss';
import { ChatContextProvider } from '../../contexts/chat';
import Main from '../Main';

const App = (): JSX.Element => (
  <ChatContextProvider>
    <Main />
  </ChatContextProvider>
);

export default App;
