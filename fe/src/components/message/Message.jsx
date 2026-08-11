import React from 'react';
import PropTypes from 'prop-types';

export const Message = ({ 
  children, 
  type = 'info', 
  className = '', 
  icon,
  dismissible = false,
  onDismiss,
  show = true
}) => {
  if (!show) return null;

  const typeStyles = {
    info: 'bg-[#2a2b38] border-blue-500/30 text-blue-300',
    success: 'bg-[#2a2b38] border-green-500/30 text-green-300',
    warning: 'bg-[#2a2b38] border-yellow-500/30 text-[#ffeba7]',
    error: 'bg-[#2a2b38] border-[#ff6384]/30 text-[#ff6384]',
    neutral: 'bg-[#2a2b38] border-gray-600/30 text-[#c4c4c4]'
  };

  const iconMap = {
    info: 'ℹ️',
    success: '✅',
    warning: '⚠️',
    error: '❌',
    neutral: '📄'
  };

  const baseClasses = `
    relative p-4 mb-4 border rounded-lg
    ${typeStyles[type] || typeStyles.neutral}
    ${className}
  `.trim().replace(/\s+/g, ' ');

  return (
    <div className={baseClasses} role="alert">
      <div className="flex items-start">
        {icon && (
          <span className="mr-3 text-lg" aria-hidden="true">
            {typeof icon === 'string' ? icon : iconMap[type]}
          </span>
        )}
        <div className="flex-1">
          {children}
        </div>
        {dismissible && onDismiss && (
          <button
            onClick={onDismiss}
            className="ml-3 text-lg leading-none text-[#c4c4c4] hover:text-white focus:outline-none"
            aria-label="Dismiss message"
          >
            ×
          </button>
        )}
      </div>
    </div>
  );
};

Message.propTypes = {
  children: PropTypes.node.isRequired,
  type: PropTypes.oneOf(['info', 'success', 'warning', 'error', 'neutral']),
  className: PropTypes.string,
  icon: PropTypes.oneOfType([PropTypes.string, PropTypes.node]),
  dismissible: PropTypes.bool,
  onDismiss: PropTypes.func,
  show: PropTypes.bool
};

export default Message;