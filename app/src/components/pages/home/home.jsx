import React from "react";
import { connect } from 'react-redux';
import { Button } from '@material-ui/core';

import { FileSelector } from '../../shared/fileSelector';
import { transformData } from '../../../store/actions';

import style  from './home.module.css';

export class HomeComponent extends React.Component {
  title = 'Trello Transform'
  subTitle = 'Transform json file exported from trello to useful information'
  typeArray = ['json', 'csv']

  object = {
    list: ['cherie', 'cherie2', 'cherie3'],
    label: {
      network: 2,
      book: 3
    },
    test: 'test'
  }

  fileUploadHandler = (file) => {
    this.props.transformData(file)
  }

  render() {
    return (
      <div className={`${style.homeContainer} animated fadeInDown slow`}>
        <div>
          <h1>{this.title}</h1>
          <span>{this.subTitle}</span>
          <div className={style.uploadFile}>
            <FileSelector uploadFile={this.fileUploadHandler} typeArray={this.typeArray}></FileSelector>
            <div style={{flex:1}}></div>
            <Button variant="outlined">SUBMIT</Button>
          </div>
        </div>
      </div>
    )
  }
}

export const Home = connect(null, {transformData})(HomeComponent)