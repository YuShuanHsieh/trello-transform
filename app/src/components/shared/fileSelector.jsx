import React from 'react';
import PropTypes from 'prop-types';

import style from './fileSelector.module.css';
import { isValidFileType } from '../../utilities/validators';

export class FileSelector extends React.Component {
  static propTypes = {
    validTypes: PropTypes.arrayOf(PropTypes.string),
    uploadFile: PropTypes.func.isRequired,
  }

  static defaultProps = {
    validTypes: [],
  }

  fileRef = React.createRef()

  state = {
    label: `please select a ${this.props.validTypes} file`,
  }

  handleChangeFile = () => {
    if (!isValidFileType(this.fileRef.current.files, ...this.props.validTypes)) {
      this.setState({
        label: `Invalid type. Please select a file with ${this.props.validTypes} type`,
      });
      return;
    }
    this.setState({
      label: this.fileRef.current.files[0].name,
    });
    this.props.uploadFile(this.fileRef.current.files[0]);
  }

  render() {
    return (
      <React.Fragment>
        <input type="file" id="fileElem" className={style.visuallyHidden} ref={this.fileRef} onChange={this.handleChangeFile} multiple />
        <label htmlFor="fileElem">{this.state.label}</label>
      </React.Fragment>
    );
  }
}
