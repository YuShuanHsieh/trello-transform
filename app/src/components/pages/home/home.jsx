/* eslint-disable react/forbid-prop-types */
import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { Button } from '@material-ui/core';

import { FileSelector } from '../../shared/fileSelector';
import { transformData } from '../../../store/actions';
import { routePath } from '../../../configuration';

import style from './home.module.css';

export class HomeComponent extends React.Component {
  static propTypes = {
    history: PropTypes.object.isRequired,
    transformData: PropTypes.func.isRequired,
  }

  state = {
    file: '',
  }

  title = 'Trello Transform';

  subTitle = 'Transform json file exported from trello to useful information';

  validTypes = ['json'];

  handleUploadFile = (uploadFile) => {
    this.setState({ file: uploadFile });
  }

  handleSubmit = () => {
    if (this.state.file) {
      this.props.transformData(this.state.file);
      this.props.history.push(routePath.result);
    }
  }

  render() {
    return (
      <div className={`${style.homeContainer} animated fadeInDown slow`}>
        <div>
          <h1>{this.title}</h1>
          <span>{this.subTitle}</span>
          <div className={style.uploadFile}>
            <FileSelector uploadFile={this.handleUploadFile} validTypes={this.validTypes} />
            <div style={{ flex: 1 }} />
            <Button variant="outlined" disabled={!this.state.file} onClick={this.handleSubmit}>SUBMIT</Button>
          </div>
        </div>
      </div>
    );
  }
}

export const Home = withRouter(connect(null, { transformData })(HomeComponent));
