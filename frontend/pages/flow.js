import { useRouter } from 'next/router'
import { useState } from 'react'

import styles from "../styles/flow.module.css"

const defaultModel = {
    username: "mattia",
    password: "test"
}

export default function () {
    const router = useRouter()
    const [user, setUser] = useState(defaultModel)

    const onClick = () => {
        fetch("http://localhost:6969/api/login", {
            method: "POST",
            body: JSON.stringify(user),
            credentials: "include"
        });
    }

    return (
        <center>
            <pre>{JSON.stringify(user, null, 2)}</pre>
            <div className={styles.container}>
                <input className={styles.input} type="text" placeholder="Username" />
                <input className={styles.input} type="password" placeholder="Password" />
                <br />
                <button className={styles.btn} onClick={onClick}>Login</button>
            </div>
        </center>
    )
}
