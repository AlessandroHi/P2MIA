import { Link, useParams } from "react-router-dom";
import { Routes, Route, HashRouter } from 'react-router-dom'
import VerticalNav from '../../navigation/VerticalNav';
import "../../App.css"; 
import loginIcon from "../../assets/login.png";

export default function SignIn({ ip }) {
  const { disk, part } = useParams();

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("submit", disk, part);

    const user = e.target.uname.value;
    const pass = e.target.psw.value;

    // Validar los datos ingresados
    if (user === "root" && pass === "123") {
      console.log("Usuario válido");
      <Route path="/login" element={<DataPartition ip={ip}/>} />
      // Aquí podrías redirigir al usuario a la página principal o realizar otras acciones
    } else {
      console.log("Usuario o contraseña incorrectos");
      // Aquí podrías mostrar un mensaje de error al usuario
    }
  };

  return (
    <div className="sign-in-container">
      <VerticalNav />
      <form onSubmit={handleSubmit} className="sign-in-form">
        <div className="form-header">
          <h2>Welcome</h2>
          <img src={loginIcon} alt="disk" style={{ width: "100px", textAlign: "center" }} />
          <p>Please sign in to continue.</p>
        </div>
        <div className="form-inputs">
          <div className="input-group">
            <label htmlFor="uname">Username</label>
            <input type="text" placeholder="Enter your username" name="uname" required />
          </div>
          <div className="input-group">
            <label htmlFor="psw">Password</label>
            <input type="password" placeholder="Enter your password" name="psw" required />
          </div>
        </div>
        <div className="form-actions">
          <button type="submit">Sign In</button>
        </div>
      </form>
    </div>
  );
}