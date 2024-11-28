import axios from "axios"
import { useState } from "react"
import { Button, Stack } from "react-bootstrap"
import ApiConfig from "../ApiConfig"

const ConnectInfo = () => {
    
    const [server, setServer] = useState("")
    const [clicks, setClicks] = useState(0)
    const [timeUsed, setTimeUsed] = useState()

    const btnClicked = async () => {
        const start_dte = new Date().getTime()
        await axios.get(ApiConfig.loc).then((response) => {setServer(response.data.server_name); setClicks(response.data.click_count);}).catch(err => alert(err))
        const end_dte = new Date().getTime()
        setTimeUsed(end_dte - start_dte)
    }
    
    return (
        <Stack>
            <Button onClick={btnClicked}>Запрос</Button>
            <div style={{marginBlockStart: "50px"}}></div>
            {
                server !== "" && (
                    <>
                        <p>Текущий сервер: {server}</p>
                        <p>Кол-во кликов: {clicks}</p>
                        <p>Время запроса (мс): {timeUsed}</p>       
                    </>
                )
            }
        </Stack>
    )
}

export default ConnectInfo