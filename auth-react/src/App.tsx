import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from "axios";
import { getEnvironmentData } from 'worker_threads';
import { Router } from './router/Router';

const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';

export const ApiGet = (URL:string) => {
  axios.get(URL)
  .then((res) => {
    console.log(res);
  }).catch((error) => {
    
    console.log('通信失敗');
    console.log(error.status);
  });
};

function App() {
  // ApiGet(baseURL);
  return (
    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.tsx</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React
    //     </a>
    //   </header>
    // </div>
    <>
      <Router />
    </>
    
  );
}

export default App;
