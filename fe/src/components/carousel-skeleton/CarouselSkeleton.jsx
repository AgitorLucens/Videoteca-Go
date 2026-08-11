import React from 'react';
import PropTypes from 'prop-types';

export const CarouselSkeleton = ({ 
  items = 6, 
  showArrows = true, 
  showDots = true,
  className = '',
  height = 'h-64',
  centerMode = true,
  slidesToShow = 3
}) => {
  const shimmerClasses = 'bg-gradient-to-r from-gray-200 via-gray-300 to-gray-200 bg-[length:200%_100%] animate-shimmer';

  return (
    <div className={`slider-container ${className}`}>
      <div className={`relative ${height} overflow-hidden`}>
        {/* Skeleton items matching react-slick layout */}
        <div className="flex h-full">
          {Array.from({ length: items }).map((_, index) => {
            // Calculate opacity based on position (center mode effect)
            let opacity = 'opacity-40';
            if (centerMode) {
              if (index === 1 || index === 2 || index === 3) {
                opacity = index === 2 ? 'opacity-100' : 'opacity-70';
              }
            } else {
              opacity = index < slidesToShow ? 'opacity-100' : 'opacity-40';
            }

            return (
              <div
                key={index}
                className={`flex-shrink-0 w-full px-12 ${opacity}`}
                style={{
                  animation: 'shimmer 1.5s infinite',
                  animationDelay: `${index * 0.1}s`
                }}
              >
                <div className={`h-full ${shimmerClasses} rounded-lg flex items-center justify-center`}>
                  {/* Skeleton content matching carousel item structure */}
                  <div className="text-center">
                    <div className="w-16 h-16 mx-auto mb-4 rounded-lg bg-gray-300 shimmer"></div>
                    <div className="h-8 w-8 mx-auto rounded bg-gray-300 shimmer"></div>
                  </div>
                </div>
              </div>
            );
          })}
        </div>

        {/* Navigation arrows */}
        {showArrows && (
          <>
            <div className="absolute left-4 top-1/2 transform -translate-y-1/2 w-10 h-10 rounded-full bg-white/80 backdrop-blur-sm shimmer z-10"></div>
            <div className="absolute right-4 top-1/2 transform -translate-y-1/2 w-10 h-10 rounded-full bg-white/80 backdrop-blur-sm shimmer z-10"></div>
          </>
        )}
      </div>

      {/* Dots indicator */}
      {showDots && (
        <div className="flex justify-center space-x-2 mt-4">
          {Array.from({ length: Math.min(items, 6) }).map((_, index) => (
            <div
              key={index}
              className={`w-2 h-2 rounded-full ${shimmerClasses}`}
              style={{
                animation: 'shimmer 1.5s infinite',
                animationDelay: `${index * 0.2}s`
              }}
            ></div>
          ))}
        </div>
      )}
    </div>
  );
};

CarouselSkeleton.propTypes = {
  items: PropTypes.number,
  showArrows: PropTypes.bool,
  showDots: PropTypes.bool,
  className: PropTypes.string,
  height: PropTypes.string,
  centerMode: PropTypes.bool,
  slidesToShow: PropTypes.number
};

export default CarouselSkeleton;