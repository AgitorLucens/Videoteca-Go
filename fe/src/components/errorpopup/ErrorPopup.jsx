
export default function ErrorPopup({ message, onClose }) {
  // dont show if no error
  if (!message) return null; 

   return (
    <div className="fixed inset-0 flex items-center justify-center z-50">
      {/* Popup box */}
  <div className="bg-[#2a2b38] border border-[#ff6384] rounded-xl shadow-lg p-6 max-w-sm w-full">
      <h2 className="text-[#ff6384] text-lg font-bold mb-2">Error</h2>
      <p className="text-[#c4c4c4]">{message}</p>
      <button
        onClick={onClose}
        className="mt-4 px-4 py-2 bg-[#ff6384] text-white rounded-lg hover:bg-red-500 transition-colors"
      >
        Close
      </button>
    </div>
    </div>
  );
}
