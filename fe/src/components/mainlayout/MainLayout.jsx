import { Outlet } from 'react-router-dom';
import Header from "../header/Header.jsx"; // Adjust path as needed

export default function MainLayout() {
  return (
    <div>
      <Header />
      
      <main>
        <Outlet /> 
      </main>
    </div>
  );
}