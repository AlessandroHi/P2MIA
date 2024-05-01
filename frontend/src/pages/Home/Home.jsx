
import VerticalNav from '../../navigation/VerticalNav'
import ComandConsole from "../ComandConsole/ComandConsole";

export default function Home({ip}) {
  return (
    <>
      <VerticalNav/>
      <ComandConsole ip={ip}/>
   </>
  )
}