import React, { useState, useCallback } from 'react';
import { useSocketIO } from '../../hooks/useSocketIO';

function App(): JSX.Element {
  const [messageInput, setMessageInput] = useState('');
  const { isConnected, messages, sendMessage } = useSocketIO(process.env.SERVER_URL);

  const handleChange = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;

    setMessageInput(value);
  }, []);

  const handleSend = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
    sendMessage(messageInput);
    setMessageInput('');
  };

  return (
    <main>
      <h3>{isConnected ? 'Connected' : 'Not Connected'}</h3>
      <div>
        <ol>
          {messages.map((message: any) => (
            <li>{JSON.stringify(message)}</li>
          ))}
        </ol>
        <form onSubmit={handleSend}>
          <input type="text" value={messageInput} onChange={handleChange} />
          <button>Send</button>
        </form>
      </div>
    </main>
  );
};

export default App;
