import partitionIMG from "../../assets/partition.png";
import { useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import VerticalNav from '../../navigation/VerticalNav'

export default function Partition({ip="localhost"}) {
  const { id } = useParams() //LETRA DISCO
  const [data, setData] = useState([])
  const navigate = useNavigate()
  const [data2, setData2] = useState([])


  useState(() => {
 
    var dataF = {
      User: 'root',
      Password: 'admin'
    }
    console.log(`fech to http://${ip}:4000/`)
    fetch(`http://${ip}:4000/partition`, {
      method: 'POST', 
      headers: {
        'Content-Type': 'application/json' 
      },
      body: JSON.stringify({ User: id })
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
    console.log("click", objIterable)
    navigate(`/login/${id}/${objIterable}`)
  }

  return (
    <>

      <VerticalNav/>

     {/*  <h1>{data2.Status}</h1>
      <h2>{data2.Value}</h2> */}

      <ddiv style={{
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
                <img src={partitionIMG} alt="disk" style={{ width: "100px" }} />
                <p>{objIterable}</p>
              </div>
            )
          })
        }

      </ddiv>
    </>
  )
}