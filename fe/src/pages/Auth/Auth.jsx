import { useState } from "react";
import CreateUserCard from "../../components/createusercard/CreateUserCard";
import LoginCard from "../../components/logincard/LoginCard";
import AuthToggle from "../../components/authtoggle/AuthToggle";
import {IconYinYang} from "../../components/sprite/Sprite";

export default function Auth() {
  const [isFlipped, setIsFlipped] = useState(false);

  return (
    <div className="flex flex-col items-center">
      {/* Toggle Button */}
      <div className="uppercase mt-4 sm:mt-0"> 
        <h6 className="mb-0 pb-3 text-sm sm:text-base"><span>Sign In </span><span>Sign Up</span></h6>
      </div>

      <AuthToggle  checked={isFlipped}
                  onChange={setIsFlipped}
                  className="mt-4"/>
      
      {/* Flip Container */}
      <div className="relative w-full max-w-[400px] px-4 h-[500px] [perspective:1000px]">
        <div
          className={`relative w-full h-full transition-transform duration-500 [transform-style:preserve-3d] ${
            isFlipped ? "[transform:rotateY(180deg)]" : ""
          }`}
        >
          {/* Front Side - Login */}
          <div className="absolute w-full h-full  backface-hidden">
            <LoginCard />
          </div>

          {/* Back Side - Create User */}
          <div className="absolute w-full h-full [transform:rotateY(180deg)] backface-hidden">
            <CreateUserCard />
          </div>
        </div>
      </div>

      
    </div>
  );
}
