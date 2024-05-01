import { Link, useParams } from "react-router-dom";
import VerticalNav from '../../navigation/VerticalNav';
import "../../App.css"; 
import loginIcon from "../../assets/login.png";
export default function SignIn({ ip = "localhost" }) {
   const { disk, part } = useParams();
 
   const handleSubmit = (e) => {
     e.preventDefault();
     console.log("submit", disk, part);
 
     const user = e.target.uname.value;
     const pass = e.target.psw.value;
 
     console.log("user", user, pass);
   };
 
   return (
     <div className="sign-in-container"> {/* Contenedor principal */}
       <VerticalNav />
       <form onSubmit={handleSubmit} className="sign-in-form"> {/* Formulario de inicio de sesión */}
         <div className="form-header"> {/* Encabezado del formulario */}
           <h2>Welcome</h2>
           <img src={loginIcon} alt="disk" style={{width: "100px",  textAlign: "center"}} />
           <p>Please sign in to continue.</p>
         </div>
         <div className="form-inputs"> {/* Campos de entrada del formulario */}
           <div className="input-group">
             <label htmlFor="uname">Username</label>
             <input type="text" placeholder="Enter your username" name="uname" required />
           </div>
           <div className="input-group">
             <label htmlFor="psw">Password</label>
             <input type="password" placeholder="Enter your password" name="psw" required />
           </div>
         </div>
         <div className="form-actions"> {/* Acciones del formulario */}
           <button type="submit">Sign In</button>
         </div>
       </form>
     </div>
   );
 }