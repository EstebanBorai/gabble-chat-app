import React, { useRef } from 'react';
import Main from '../Main';
import { ChatContextProvider } from '../../contexts/chat';
import { UserContextProvider } from '../../contexts/user';

const App = (): JSX.Element => (
  <ChatContextProvider>
    <UserContextProvider>
      <Main />
    </UserContextProvider>
  </ChatContextProvider>
);

export default App;
