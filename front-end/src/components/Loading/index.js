import React from 'react';
import { Progress } from 'reactstrap';

import styles from './styles.module.css'; 

const Loading = (props) => {
  return (
    <div>
      <Progress
        className={ styles.Progress }
        animated striped value={ props.progress }
      />
      <p 
        className={ styles.Paragraph }
      >
        Getting the repositories list from Github...
      </p>
    </div>
  );
}

export default Loading;