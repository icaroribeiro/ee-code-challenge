import React, { Component } from 'react';
import BootstrapTable from 'react-bootstrap-table-next';
import ToolkitProvider, { Search } from 'react-bootstrap-table2-toolkit';

import Tag from './Tag';
import TagModal from './TagModal';

import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

class Repository extends Component {
  render() {
    const { SearchBar } = Search;

    const username = this.props.username;
    const data = this.props.data;

    const tagsFormatter = (cell, row, rowIndex, formatExtraData) => {
      return (
        <Tag 
        repository={ row }
        />
      );
    }

    const modalFormatter = (cell, row, rowIndex, formatExtraData) => {
      return (
        <TagModal 
          buttonLabel={ "edit" }
          username={ username }
          repository={ row }
        />
      );
    }

    const columns = [
      { 
        dataField: "id",
        text: "Id",
        hidden: true,
        searchable: false
      },
      {
        dataField: "name",
        text: "Repository",
        align: "left",
        searchable: false
      },
      { 
        dataField: "description",
        text: "Description",
        align: "left",
        searchable: false
      },
      { 
        dataField: "language",
        text: "Language",
        align: "left",
        searchable: false
      },
      { 
        dataField: "tags",
        text: "Tags",
        align: "left",
        formatter: tagsFormatter
      },
      { 
        dataField: "edit",
        text: "",
        align: "left",
        searchable: false,
        formatter: modalFormatter
      }
    ];

    function onColumnMatch ({searchText, value, column, row}) {
      if (searchText === null || searchText === "") {
        return true;
      } else {
        const tags = value;

        if (tags) {
          for (var i = 0; i < tags.length; i++) {
            if ((tags[i]).startsWith(searchText)) {
              return true;
            }
          }
        }
    
        return false;
      }
    }

    return (      
      <ToolkitProvider
      keyField="id"
      data={ data }
      columns={ columns }
      search={ {
        onColumnMatch
      } }
      bootstrap4={ true }
      >
      {
        props => (
          <div style={ { padding: "0px", margin: "20px" } }>
            <SearchBar placeholder="search by tag" { ...props.searchProps } />
            <BootstrapTable
              { ...props.baseProps }
            />
          </div>
        )
      }
      </ToolkitProvider>
    );
  }
}

export default Repository;