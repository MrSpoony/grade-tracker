import Head from 'next/head'
import styles from '../styles/App.module.css'

export default function App() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Grades</title>
      </Head>
    </div>
  )
}
