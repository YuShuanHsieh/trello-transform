import React, { Component } from 'react';
import { connect } from 'react-redux';

import { APPList } from './components/list';
import { transformData } from './store/actions';
import { isValidFileType } from './utilities/validators';
import './App.css';

class App extends Component {
  fileRef = React.createRef()

  changeFileHandler = () => {
    if (!isValidFileType(this.fileRef.current.files, 'json')) return;
    const formData = new FormData();
    formData.append('file', this.fileRef.current.files[0])
    this.props.transformData(formData);
  }



  render() {
    return (
      <React.Fragment>
        <input type="file" ref={this.fileRef} onChange={this.changeFileHandler}/>
        <APPList list={this.props.list}/>
      </React.Fragment>
    );
  }
}

const mapStateToProps = state => ({
  list: state.list
})

export const AppContainer = connect(mapStateToProps, {transformData})(App)