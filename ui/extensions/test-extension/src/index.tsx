import * as React from 'react'
import styles from './styles.module.css'
export * from './components/Component1'

interface Props {
  text: string
}

export const ExampleComponent = ({ text }: Props) => {
  return <div className={styles.test}>Example Component: {text}</div>
}
