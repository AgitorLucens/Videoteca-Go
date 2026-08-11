import { IconFilm } from '../sprite/Sprite.jsx';

export default function LoadingSpinner({className}) {
  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black/40 backdrop-blur-sm z-50">
        <IconFilm className={className}/>
    </div>
  );
}
