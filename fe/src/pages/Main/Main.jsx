import { Outlet } from 'react-router-dom';

const Main = () => {

	return ( 
        <div>

            {/* Outlet renderiza componentes de una ruta padre, en este caso "Main" */}
            <Outlet />
        </div>
	);
};

export default Main;