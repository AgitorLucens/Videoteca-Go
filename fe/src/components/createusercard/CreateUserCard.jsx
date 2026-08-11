import { useState } from "react";
import { createUserService } from "../../services/createuser.service"
import InputIcon from "../inputicon/InputIcon.jsx";
import { IconArroa, IconUser, IconLock } from "../sprite/Sprite.jsx";
import ErrorPopup from "../errorpopup/ErrorPopup.jsx";

export default function CreateUserCard() {
    //Login Inputs
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [passwordconfirm, setPasswordConfirm] = useState("");
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");

    async function handleSubmit(e) {
        e.preventDefault();
        setError("");
        setSuccess("");

        if (!email || !password) {
            setError("Email and password are required");
            return;
        }
        try {
            const response = await createUserService({ name, email, username, password, passwordconfirm,"role": "user" });
            setSuccess("Account created successfully! You can now log in.");
            setName("");
            setEmail("");
            setUsername("");
            setPassword("");
            setPasswordConfirm("");
        } catch (error) {
            setError(error.response?.data?.error || "Failed to create account");
        }
    };

    return (
        <div className="w-full h-full bg-[#2a2b38] rounded-lg p-6 flex flex-col justify-center items-center ">
            <h4 className="text-2xl font-semibold mb-4">Create a New Account</h4>
            <ErrorPopup message={error} onClose={() => setError("")} />
            {success && <p className="text-green-400 text-sm mb-2 text-center">{success}</p>}
            <form className="w-full flex flex-col gap-4" onSubmit={handleSubmit}>
                <InputIcon  type="text"
                            placeholder="Your Full Name"
                            icon={IconUser}
                            onChange={(e) => setName(e.target.value)}/>      
                <InputIcon  type="email"
                            placeholder="Your Email"
                            icon={IconArroa}
                            onChange={(e) => setEmail(e.target.value)}/>  
                <InputIcon  type="text"
                            placeholder="Your Username"
                            icon={IconUser}
                            onChange={(e) => setUsername(e.target.value)}/>
                <InputIcon  type="password"
                            placeholder="Your Password"
                            icon={IconLock}
                            onChange={(e) => setPassword(e.target.value)}/>
                <InputIcon  type="password"
                            placeholder="Your Password"
                            icon={IconLock}
                            onChange={(e) => setPasswordConfirm(e.target.value)}/>                                        
                <button
                    type="submit"
                    className="bg-[#ffeba7] text-[#102770] 
                         hover:bg-[#102770] hover:text-[#ffeba7] 
                                font-semibold py-2 rounded-md 
                                h-[44px] w-[118px]
                                mx-auto
                                text-transform: uppercase"
                >
                    Register
                </button>
            </form>
        </div>
    );

}