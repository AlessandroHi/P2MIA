import { Link } from "react-router-dom";
import "../App.css"; 

export default function VerticalNav() {
  return (
    <nav className="vertical-nav">
      <ul className="nav-list">
        <li className="nav-item">
          <Link to="/" className="nav-link">Pantalla1</Link>
        </li>
        <li className="nav-item">
          <Link to="/DiskCreen" className="nav-link">Pantalla2</Link>
        </li>

        <li className="nav-item">
          <Link to="/Reports" className="nav-link">Pantalla3</Link>
        </li>
        {/* Agrega más elementos de lista y enlaces según sea necesario */}
      </ul>
    </nav>
  );
}
