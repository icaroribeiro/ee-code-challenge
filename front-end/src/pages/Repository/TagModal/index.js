import React from 'react';
import { Button, Modal, ModalHeader, ModalBody, Form, FormGroup, Input, ModalFooter } from 'reactstrap';

import { connect } from 'react-redux';
import { Container } from 'reactstrap';

import styles from './styles.module.css'; 
import { getRepository } from '../../../service/api.js';
import { updateUserRepository } from '../../../service/api.js';

import { renewRepository } from '../../../actions/repository.js';

function createStringWithTags(tags) {
  var str = "";

  if (tags) {
    for (var i = 0; i < tags.length; i++) {
      str = str + tags[i];

      if (i < tags.length - 1) {
        str = str + ", ";
      }
    }
  }

  return str;
}

function splitStringIntoArray(str) {
  var array = [];

  if (str === "") {
    return array;
  }
  
  var newStr = str.replace(/[\s,]+/g, ' ').trim();
  array = newStr.split(" ");
  let removeDups = (list) => list.filter((v, i) => list.indexOf(v) === i);  

  return removeDups(array);
}

function createNewTags(tags) {  
  var newTags = [];
  
  if (tags.length > 0) {
    for (var i = 0; i < tags.length; i++) {
      newTags.push(tags[i]);
    }
  }

  return newTags;
}

function hasToUpdateTags(newTags, tags) {
  var hasToUpdate = false;

  if ((typeof tags === 'undefined') || (tags && tags.length === 0)) {
    if (newTags.length !== 0) {
      hasToUpdate = true;
    }
  } else {
    if (tags.length !== newTags.length) {
      hasToUpdate = true;
    }

    for (let i = 0; i < newTags.length; i++) {
      var pos = tags.findIndex(tag => tag === newTags[i]);
    
      if (pos !== -1) {
        newTags[i] = tags[pos];
      } else {  
        hasToUpdate = true;
      }
    }
  }

  return hasToUpdate;
}

class TagModal extends React.Component {
  constructor(props) {
    super(props);
    
    this.state = {
      modal: false,
      tags: [],
      tagsString: ""
    };

    this.toggle = this.toggle.bind(this);
    this.save = this.save.bind(this);    
    this.cancel = this.cancel.bind(this);
  }

  toggle() {
    this.setState(prevState => ({
      modal: !prevState.modal
    }));
  }

  save() {
    var tags = splitStringIntoArray(this.state.tagsString);
    var newTags = createNewTags(tags);
    var repository = this.props.repository;
    var username = this.props.username;

    if (hasToUpdateTags(newTags, repository.tags)) {
      repository.tags = newTags;
      this.props.onRenewRepository(username, repository.id, repository)
          .then(() => {
            if (newTags.length === 0) {  
              if ((typeof repository.tags === 'undefined') || (repository.tags && repository.tags.length === 0)) {
                  this.loadSuggestedTags(repository);
                }
            } else {
              this.setState({
                tags: newTags,
                tagsString: createStringWithTags(newTags)
              });
            }
        })
      .catch((err) => {
        console.log('Caught error: ', err);
      })
    }

    this.setState(prevState => ({
      modal: !prevState.modal
    }));
  };

  cancel() {
    var tags = this.state.tags;

    this.setState(prevState => ({
      modal: !prevState.modal,
      tagsString: createStringWithTags(tags)
    }));
  };

  async loadSuggestedTags (repository) {
    try {
      var suggestedTags = [];
      const response = await getRepository(repository.id);
      
      if (response.data.tags && response.data.tags.length > 0) {
        var maxTagsNbr = process.env.REACT_APP_MAX_TAGS_NBR;

        if ((typeof maxTagsNbr !== 'undefined') && (maxTagsNbr && maxTagsNbr > 0)) {
            suggestedTags = response.data.tags.slice(0, maxTagsNbr);
        }
      }

      this.setState({
        tags: suggestedTags,
        tagsString: createStringWithTags(suggestedTags)
      });
    }
    catch (error) {
      console.log(error);
    }
  }

  componentDidMount() {
    var repository = this.props.repository;
  
    if ((typeof repository.tags === 'undefined') || (repository.tags && repository.tags.length === 0)) {
        this.loadSuggestedTags(repository);
    } else {
      this.setState({
        tags: repository.tags,
        tagsString: createStringWithTags(repository.tags)
      });
    }
  }

  handleChange = event => {
    this.setState({
      tagsString: event.target.value
    });
  }

  render() {
    return (
      <div>
        <Button color="link" onClick={ this.toggle }>{ this.props.buttonLabel }</Button>
        <Container className={ styles.Container }>
          <Modal isOpen={ this.state.modal } toggle={ this.toggle }>
            <ModalHeader toggle={ this.cancel }>edit tags for { this.props.repository.name }</ModalHeader>
            <ModalBody>
              <Form>
                <FormGroup>
                  <Input name="tagsString" value={ this.state.tagsString } onChange={ this.handleChange }/>
                </FormGroup>
              </Form>
            </ModalBody>
            <ModalFooter style={ { alignItems: "center", justifyContent: "center" } }>
              <Button color="primary" onClick={ this.save }>Save</Button>
              <Button color="secondary" onClick={ this.cancel }>Cancel</Button>
            </ModalFooter>
          </Modal>
        </Container>
      </div>
    );
  }
}

const mapDispatchToProps = dispatch => {
  return {
    onRenewRepository: async (username, id, repository) => {
      try {
        const response = await updateUserRepository(username, id, repository);
        dispatch(renewRepository(response.data));
      }
      catch (err) {
        console.log('Caught error: ', err);
      }
    }
  };
};

export default connect(null, mapDispatchToProps)(TagModal);