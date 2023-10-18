import { Route, Routes } from "react-router-dom";
import {Login} from "../components/Login";
import {Index} from "../components/Index";

export const Router = () => {
    return (
        <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/login" element={<Login />} />
        </Routes>
    )
};