import { useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";
import "../../App.css"; 

export default function ComandConsole({ip="localhost"}) {
  const { id } = useParams()
  const [inputValue, setInputValue] = useState("");
  const [outputValue, setOutputValue] = useState("");

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleEnterPress = (event) => {
    if (event.key === "Enter") {
      executeCommand();
    }
  };

  const executeCommand = () => {
    const command = inputValue.trim();
    fetch(`http://${ip}:4000/comand`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ User: command }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        if (Array.isArray(data)) {
          setOutputValue(data.join("\n"));
        } else {
          setOutputValue(data.Value);
        }
      })
      .catch((error) => {
        console.error("Error:", error);
        setOutputValue("Error al procesar la solicitud");
      });

    setInputValue("");
  };

  return (
    <div className="command-console-container">
      <div className="command-console">
        <textarea
          className="output-area"
          value={outputValue}
          readOnly
          rows={5}
          style={{ resize: "none" ,width: "99%", height: "90%", marginBottom: "12px"}}
        />
        <input
          type="text"
          className="input-area"
          value={inputValue}
          onChange={handleInputChange}
          onKeyDown={handleEnterPress}
          placeholder="Enter command..."
          style={{ width: "100%" , borderRadius: "15px" , height: "45px"}}
        />
      </div>
    </div>
  );
}
