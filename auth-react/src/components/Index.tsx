
import logo from './img/logo.svg';
import './css/Index.css'
import { Link } from "react-router-dom";

export const Index = () =>{
    return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <Link to="/sandbox" className='App-link'>Sandbox Page</Link>
          </header>
        </div>
    )
};

