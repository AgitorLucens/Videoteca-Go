# Carousel Skeleton Component

A shimmer skeleton loading component designed specifically for carousel placeholders.

## Features

- **Shimmer Animation**: Smooth gradient animation that mimics content loading
- **Customizable Items**: Configure the number of skeleton items to display
- **Navigation Elements**: Optional skeleton arrows and dots indicators
- **Responsive Design**: Uses Tailwind CSS classes for consistent styling
- **Staggered Animation**: Each element animates with a slight delay for a wave effect

## Usage

```jsx
import { CarouselSkeleton } from '@/components/carousel-skeleton';

// Basic usage
<CarouselSkeleton />

// Custom configuration
<CarouselSkeleton 
  items={3}
  showArrows={true}
  showDots={true}
  height="h-80"
  itemHeight="h-60"
  className="my-4"
/>

// Minimal skeleton (no arrows or dots)
<CarouselSkeleton 
  items={4}
  showArrows={false}
  showDots={false}
/>
```

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `items` | number | 5 | Number of skeleton items to display |
| `showArrows` | boolean | true | Show skeleton navigation arrows |
| `showDots` | boolean | true | Show skeleton dots indicator |
| `className` | string | '' | Additional CSS classes |
| `height` | string | 'h-64' | Height of the carousel container |
| `itemHeight` | string | 'h-48' | Height of individual skeleton items |

## Styling

The component uses Tailwind CSS classes and includes custom CSS animations defined in `shimmer.css`. Make sure to import the CSS file:

```jsx
import '@/components/carousel-skeleton/shimmer.css';
```

## Animation Details

- **Shimmer Effect**: Gradient moves from left to right continuously
- **Staggered Delays**: Each item starts animation 100ms after the previous
- **Duration**: 1.5 seconds per animation cycle
- **Easing**: Smooth ease-in-out timing function

## Accessibility

- Uses `aria-hidden="true"` for decorative elements
- Semantic HTML structure for screen readers
- High contrast skeleton colors for better visibility