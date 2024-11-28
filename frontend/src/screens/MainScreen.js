import {Container } from "react-bootstrap"
import ConnectInfo from "../components/ConnectInfo"

const MainScreen = () => {
    return (
        <>
            <Container style={{textAlign: "center", marginBlockStart: "10em"}}>
                <h1>Балансировщик</h1>
                <ConnectInfo />
            </Container>
        </>
    )
}

export default MainScreen