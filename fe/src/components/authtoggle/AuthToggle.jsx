import { useState } from "react";
import { IconYinYang } from "../../components/sprite/Sprite";

export default function AuthToggle({
                                checked = false,
                                onChange,
                                className = "",
                                ariaLabel = "Toggle",
                        }) {
  return (
    <button
      type="button"
      aria-pressed={checked}
      aria-label={ariaLabel}
      onClick={() => onChange?.(!checked)}
      className={`relative inline-flex h-8 w-12 items-center justify-center mx-auto ${className}`}
    >
      <span className="absolute top-1/2 -translate-y-1/2 h-3 w-full rounded-full bg-[#ffeba7] shadow-inner" />

      <span
        className={`absolute top-[-0.1px] left-[-10px] h-7 w-7 rounded-full shadow-md flex items-center justify-center
        bg-gradient-to-br from-[#0b3a7e] to-[#102770]
        transition-transform duration-300 ease-out
        ${checked ? "translate-x-10" : "translate-x-0"}`}
      >
        <IconYinYang
          className={`w-3 h-3 transition-transform duration-300 ease-out ${checked ? "rotate-180" : ""}`}
          leftColor="#0b3a7e"
          rightColor="#ffeba7"
        />
      </span>
    </button>
  );
}