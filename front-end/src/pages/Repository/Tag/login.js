import React from 'react';

const Tags = (props) => {
  var tags = props.repository.tags;

  return (
    <div>
      {
        tags && 
          tags.map(tag => {
            return (` #${tag}`);
          })
      }
    </div>
  );
};

export default Tags;