import axios from "axios";
import { error } from "console";
import { stat } from "fs";
import { useState,useContext, useEffect, useReducer } from "react";

export const PostLogin = (name: string, password: string) => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    
    axios.post(baseURL + 'login', {"name": name, "password": password})
    .then((res)=> {
        console.log(res);
        const json = res.data;
        localStorage.setItem('token', json.data.token);
        axios.defaults.headers.common["Authorization"] = "Bearer " + json.data.token;
        return true;
    }).catch((error) => {
        console.log('通信失敗');
        console.log(error.status);
        return false;
    });
};
export const PostLogin2 = async (name: string, password: string) => {
    
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    return await axios.post(baseURL + 'login', {"name": name, "password": password})
}

export const GetLogin = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    const token = localStorage.getItem('token');
    if (token === null){
        console.log('token error');
        return false;
    }
    axios.defaults.headers.common["Authorization"] = "Bearer " + token;
    axios.get(baseURL + 'login', )
    .then((res)=>{
        const json = res.data;
        console.log(res);
        localStorage.setItem('token', json.data.token);
        sessionStorage.setItem('user', json.data.name);
        sessionStorage.setItem('role', json.data.role);
    }).catch((error) =>{
        console.log('通信失敗');
        console.log(error.status);
        localStorage.removeItem('token');
    });
}

export default function GetLoginFunc() {
    return false;
};

export const PostLogout = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
       
    axios.post(baseURL + 'logout')
    .then((res)=> {
        console.log(res);
    }).catch((error) => {
        console.log('通信失敗');
        console.log(error.status);
    });
    localStorage.removeItem('token');
    sessionStorage.removeItem('user');
    sessionStorage.removeItem('role');
    delete axios.defaults.headers.common["Authorization"];

}

export const TestSend = () => {
    const baseURL:string = process.env.REACT_APP_API_URL+ '/api/v1/';
    const token = localStorage.getItem('token');
    const getLogin = async () => {
        try {
            axios.defaults.headers.common["Authorization"] = "Bearer " + token;
                const response = await axios.get(baseURL);
            console.log(response);
        }catch (error){
            console.log(error);
        }
    };
    getLogin();

}

export const Login = () => {
    const [form, setForm] = useState({
        name: '',
        password: ''
    });
    const handleForm = (e: any) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }
    const SendForm = () => {
        console.log(form);
        PostLogin(form.name, form.password);
        
    }
    const GetForm = () => {
        GetLogin();
    }
    
    const TestForm = () => {
        TestSend();
    }
    return (
        <form>
            <div>
                <label htmlFor="name">名前</label>
                <input id="name" name="name" type="text" onChange={handleForm}/>
            </div>
            <div>
                <label htmlFor="password">パスワード</label>
                <input id="password" name="password" type="password" onChange={handleForm}/>
            </div>
            <div>
                <button type="button" onClick={SendForm}>送信</button>
            </div>
            <div>
                <button type="button" onClick={GetForm}>GET</button>
            </div>
            <div>
                <button type="button" onClick={PostLogout}>Logout</button>
            </div>
            <hr />
            <div>
                <button type="button" onClick={TestForm}>TEST Send</button>
            </div>
        </form>
    )
};

export const LoginPage2 = () => {
    
    const [form, setForm] = useState({
        name: '',
        password: ''
    });
    const handleForm = (e: any) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }
    //localstorageのtoken情報が更新した場合に、再レンダリングする
    const [token, setToken] = useState(false);
    useEffect(() => {
        const tmp = localStorage.getItem('token');
        setToken(tmp !== null);
    }, [localStorage.getItem('token')]);
    const login = () => {
        //PostLoginの結果がtrueの場合は、tokenをtrueにする
        const result = PostLogin2(form.name, form.password);
        result.then((res) => {
            console.log(res);
            localStorage.setItem('token', res.data.data.token);
            setToken(true);
        }).catch((error) => {
            console.log('通信失敗');
            console.log(error.status);
        });
    }
    const logout = () => {
        PostLogout();
        setToken(false);
    }
    return (
        //tokenがnullの場合は、ログイン画面を表示する
        //tokenがnullでない場合は、ログアウト画面を表示する
        !token ? 
        <form>
        <div>
            <label htmlFor="name">名前</label>
            <input id="name" name="name" type="text" onChange={handleForm}/>
            <label htmlFor="password">パスワード</label>
            <input id="password" name="password" type="password" onChange={handleForm}/>
            <button type="button" onClick={login}>送信</button>
        </div>
        </form>:
        <div>
            <button type="button" onClick={logout}>Logout</button>
        </div>
    
    )
}

export const LoginPage = () => {
    const [form, setForm] = useState({
        name: '',
        password: ''
    });
    const handleForm = (e: any) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value
        });
    }
    return (
        <form>
        <div>
            <label htmlFor="name">名前</label>
            <input id="name" name="name" type="text" onChange={handleForm}/>
            <label htmlFor="password">パスワード</label>
            <input id="password" name="password" type="password" onChange={handleForm}/>
            <button type="button" onClick={()=>PostLogin(form.name, form.password)}>送信</button>
        </div>
        </form>
    )
}

export const LogoutPage = () => {
    return (
        <div>
            <button type="button" onClick={PostLogout}>Logout</button>
        </div>
    
    )
};