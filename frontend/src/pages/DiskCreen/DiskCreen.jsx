import diskIMG from "../../assets/disk.png";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import VerticalNav from '../../navigation/VerticalNav'

export default function DiskCreen({ip}) {
  const [data, setData] = useState([]) 
  const navigate = useNavigate()
  
  // execute the fetch command only once and when the component is loaded
  useState(() => {
 
    var dataF = {
      User: 'root',
      Password: 'admin'
    }
    console.log(`fech to http://${ip}:4000/`)
    fetch(`http://${ip}:4000/tasks`, {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json' 
      },
      body: JSON.stringify(dataF)
    })
    .then(response => response.json())
    .then(data => {
      console.log(data); // Do something with the response
      setData(data.List)
    })
    .catch(error => {
      console.error('There was an error with the fetch operation:', error);
    });
  }, [])

  const onClick = (objIterable) => {
    //e.preventDefault()
    console.log("click",objIterable)
    navigate(`/disk/${objIterable}`)
  }

  return (
    <>
      <VerticalNav/>
      <div style={{
       backgroundColor: "#464646",
       display: "flex",
       flexDirection: "row",
       position: "absolute",
       top: "10%",
       left: "22%",
       borderRadius: "15px",
       height: "80%",
       width: "75%",
       textAlign: "center"

       }}>

        {
          data.map((objIterable, index) => {
            return (
              <div key={index} style={{
                display: "flex",
                flexDirection: "column", 
                maxWidth: "100px",
                marginTop: "10px",
                height: "20px"
              }}
              onClick={() => onClick(objIterable)}
              >
                <img src={diskIMG} alt="disk" style={{width: "100px",  textAlign: "center"}} />
                <p>{objIterable}</p>
              </div>
            )
          })
        }
      
      </div>
    </>
   )
 }
 