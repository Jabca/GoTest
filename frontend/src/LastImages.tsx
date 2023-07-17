import React from 'react';
import { message } from 'antd';

import axios from 'axios';

import './LastImage.css'

interface ImagePair {
  id: number
  negative_image: string;
  positive_image: string;
}

interface ImagePairsState {
  imagePairs: ImagePair[];
  loading: boolean;
}

class LastImages extends React.Component<{}, ImagePairsState> {
  constructor(props: {}) {
    super(props);
    this.state = {
      imagePairs: [],
      loading: false,
    };
  }

  componentDidMount() {
    this.fetchImagePairs();
  }

  fetchImagePairs = async () => {
    this.setState({loading: true})

    const url = 'http://0.0.0.0:8000/get_last_images';

    axios
      .get(url)
      .then((response) => {
        console.log(response.data.data)
        const data = response.data.data;

        const newImagePairs = [] as ImagePair[];
        const tmp_image_pair = {} as ImagePair;
        
        for(var el of data){
          tmp_image_pair.id = el.ID;
          tmp_image_pair.positive_image = el.positive_image
          tmp_image_pair.negative_image = el.negative_image

          newImagePairs.push(structuredClone(tmp_image_pair))
        }

        this.setState({imagePairs: newImagePairs, loading: false})
        
        // Process the response data as needed
      })
      .catch((error) => {
        console.error('Error:', error);
        message.error('Failed to download images');
        // Handle any error that occurred during the request
      });
      
  };

  render() {
    const { imagePairs, loading } = this.state;

    return (
      <div>
        {loading && 
        <p style={{
            color: '#1890ff',
            marginTop: '5%'
          }}>
            Loading images...
        </p>}
        {imagePairs && imagePairs.map((pair: ImagePair, index: number) => (
          <div key={index}>
            <div className='pair-container'>
              <h3>ID {pair.id}</h3>
              <div className='upload-container'>
                <img 
                  src={pair.positive_image} 
                  alt={`Image ${index + 1}A`}
                  style={{
                    maxHeight: '90%',
                    maxWidth: '90%'
                }} />
                
              </div>
              <div className='upload-container'>
                <img
                  src={pair.negative_image} 
                  alt={`Image ${index + 1}B`}
                  style={{
                    maxHeight: '90%',
                    maxWidth: '90%'
                }} />
              </div>
                
            </div>
          </div>
        ))}
      </div>
    );
  }
}

export default LastImages;