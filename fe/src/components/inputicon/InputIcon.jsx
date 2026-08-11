import { IconEye, IconEyeSlash } from "../sprite/Sprite";
import { useState } from "react";


export default function InputIcon({ icon: Icon,type, ...props }) {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const isPasswordField = type === "password";
  // ternary operator only work with 2 types of values, when there more create a bug
  let finalInputType = type; 
  if (isPasswordField) {
    // If it is a password field, toggle between 'text' and 'password'
    finalInputType = isPasswordVisible ? 'text' : 'password';
  }
  
  const handleToggle = () => {
    setIsPasswordVisible((prev) => !prev);
  };
  return (
    <div className="relative w-full">
      {/* Icon */}
      <Icon className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-[#ffeba7]" />

      {/* Input */}
      <input
        {...props}
        type={finalInputType}
        className="w-full pl-10 pr-4 py-3 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors"
      />

      {/* If is password put eye icon */}
      {isPasswordField && (
        <button
          type="button" // Important: Prevents form submission
          onClick={handleToggle}
          // Positioning the button on the right side
          className="absolute right-3 top-1/2 -translate-y-1/2 p-1 text-[#c4c3ca] hover:text-[#ffeba7]"
        >
          {isPasswordVisible ? (
            <IconEyeSlash className="w-5 h-5" />
          ) : (
            <IconEye className="w-5 h-5" />
          )}
        </button>
      )}
    </div>
  );
}