import React from 'react';
import { Button, Form, FormGroup, Label, Col, Input } from 'reactstrap';

import styles from './styles.module.css'; 

const CredentialsForm = (props) => {
  return (
    <Form>
      <FormGroup row style={{marginBottom: '10px'}}>
        <Label className={styles.Label} for="exampleText" sm={6}>https://github.com/</Label>
        <Col sm={6}>
          <Input
            className={styles.Input}
            type="text"
            placeholder="username"
            required="required"
            name="username"
            value={props.username}
            onChange={props.handleChange}
          />
        </Col>
      </FormGroup>
      <FormGroup check row>
        <Button
          className={styles.Button}
          color="primary"
          onClick={props.handleSubmit}
          disabled={!props.validateForm()} >
          get repositories
        </Button>
      </FormGroup>
    </Form>
  );
};

export default CredentialsForm;