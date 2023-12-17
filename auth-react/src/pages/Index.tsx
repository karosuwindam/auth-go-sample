import logo from "../public/img/logo.svg";
import "../public/css/Index.css";
import { Link } from "react-router-dom";

export const Index = () => {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} alt="logo" className="App-logo" />
        <Link to="/home" className="App-link">
          Home Page
        </Link>
        <Link to="/page" className="App-link">
          Page
        </Link>
        <Link to="/edit" className="App-link">
          Edit
        </Link>
      </header>
    </div>
  );
};
