import { useEffect } from 'react';

function App() {
  useEffect(() => {
    console.log(window.Telegram.WebApp.bg_color);
  }, []);
  return (
    <>
      <h1>Hello</h1>
      <button>Press me</button>
    </>
  );
}

export default App;
