import { useState } from "react";
import "../../App.css"; 

export default function ComandConsole() {
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
    if (command === "clear") {
      setOutputValue("");
    } else {
      setOutputValue(outputValue + "> " + command + "\n");
    }
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