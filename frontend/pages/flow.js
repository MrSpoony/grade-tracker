import { useRouter } from 'next/router'

import styles from "../styles/flow.module.css"

export default function () {
    const router = useRouter()

    const onClick = () => {
        router.push('/')
    }

    return (
        <center>
            <div className={styles.container}>
                <input className={styles.input} type="text" placeholder="Username" />
                <input className={styles.input} type="password" placeholder="Password" />
                <br />
                <button className={styles.btn} onClick={onClick}>Login</button>
            </div>
        </center>
    )
}