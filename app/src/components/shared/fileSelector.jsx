import React from "react";

import style from './fileSelector.module.css';
import { isValidFileType } from '../../utilities/validators';

export class FileSelector extends React.Component {
  fileRef = React.createRef()

  state = {
    label: `please select a ${this.props.typeArray} file`
  }

  selectFileHandler = () => {
    if (!isValidFileType(this.fileRef.current.files, ...this.props.typeArray)) {
      this.setState({
        label: `Invalid type. Please select a file with ${this.props.typeArray} type`
      });
      return;
    };
    this.setState({
      label: this.fileRef.current.files[0].name,
    });
    this.props.uploadFile(this.fileRef.current.files[0]);
  }

  render() {
    return (
      <React.Fragment>
        <input type="file" id="fileElem" className={style.visuallyHidden} ref={this.fileRef} onChange={this.selectFileHandler} multiple/>
        <label htmlFor="fileElem">{this.state.label}</label>
      </React.Fragment>
    )
  }
}