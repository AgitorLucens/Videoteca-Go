export default function Loading({ message = "Loading..." }) {
  return (
    <div className="fixed inset-0 flex items-center justify-center bg-[#1f2029]/95 backdrop-blur-sm z-50">
      <div className="flex flex-col items-center gap-6">
        {/* Dual Spinning Circles */}
        <div className="relative w-20 h-20">
          {/* Outer spinning ring */}
          <div className="absolute inset-0 w-20 h-20 border-4 border-[#ffeba7]/20 border-t-[#ffeba7] rounded-full animate-spin"></div>
          
          {/* Inner counter-rotating ring */}
          <div 
            className="absolute inset-0 w-20 h-20 border-4 border-transparent border-r-[#102770] rounded-full animate-spin" 
            style={{ animationDuration: '1.5s', animationDirection: 'reverse' }}
          ></div>
          
          {/* Center pulsing dot */}
          <div className="absolute inset-0 flex items-center justify-center">
            <div className="w-3 h-3 bg-[#ffeba7] rounded-full animate-pulse"></div>
          </div>
        </div>
        
        {/* Loading Text with Animated Dots */}
        <div className="flex items-center gap-2">
          <span className="text-[#ffeba7] text-xl font-semibold tracking-wide">
            {message}
          </span>
          <div className="flex gap-1">
            <span 
              className="w-2 h-2 bg-[#ffeba7] rounded-full animate-pulse" 
              style={{ animationDelay: '0s' }}
            ></span>
            <span 
              className="w-2 h-2 bg-[#ffeba7] rounded-full animate-pulse" 
              style={{ animationDelay: '0.2s' }}
            ></span>
            <span 
              className="w-2 h-2 bg-[#ffeba7] rounded-full animate-pulse" 
              style={{ animationDelay: '0.4s' }}
            ></span>
          </div>
        </div>
        
        {/* Optional: Progress bar effect */}
        <div className="w-64 h-1 bg-[#2a2b38] rounded-full overflow-hidden">
          <div 
            className="h-full bg-gradient-to-r from-[#102770] via-[#ffeba7] to-[#102770] animate-pulse"
            style={{ 
              animation: 'pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite'
            }}
          ></div>
        </div>
      </div>
    </div>
  );
}
