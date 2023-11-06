import {LoginPage} from "../components/Login";
import "../components/css/Index.css"
import { useEffect, useState } from "react";
import { NotFount } from "./NotFount";
import { Link } from "react-router-dom";

export const UserType = {
    name: '',
    role: ''
}

export type User = typeof UserType[keyof typeof UserType];


export const GetRole = () => {
    const user = {
        name: sessionStorage.getItem('user') || '',
        role: sessionStorage.getItem('role') || ''
    }
    return user;
}


export const Home = () => {
    const [role,setRole] = useState(GetRole().role);
    const [user,setUser] = useState(GetRole().name);
    useEffect(() => {
        const tmpUser = GetRole();
        setUser(tmpUser.name);
        tmpUser.role === '' ? setRole('guest') : setRole(tmpUser.role);
    }, [sessionStorage.getItem('user'),sessionStorage.getItem('role')]);
    //1秒ごとにsessionStorageの値を取得する
    const [count, setCount] = useState(0);
    const [delay, setDelay] = useState(1000);
    useEffect(() => {
        const interval = setInterval(() => {
            const tmpUser = GetRole();
            if(tmpUser.name ===  user && tmpUser.role === role) {
            }else{
                setUser(tmpUser.name);
                tmpUser.role === '' ? setRole('guest') : setRole(tmpUser.role);    
            }
            setCount(count + 1);
        }, delay);
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="App">
        <header className="App-header">
        {role !== 'guest' ?
            <h1>{role} Home</h1>
        :<><NotFount /></>}

        <Link to="/" className='App-link'>Index Page</Link>
          </header>
        </div>
    )
};