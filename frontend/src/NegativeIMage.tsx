import React, { useState } from 'react';
import { Upload, Button, message } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import './NegativeImage.css';
import axios from "axios"

interface NegativeImageProps {}

interface NegativeImageState {
    imageUrl: string | null;
    uploading: boolean;
  }

class NegativeImage extends React.Component<{}, NegativeImageState> {
    constructor(props: {}) {
        super(props);
        this.state = {
          imageUrl: "",
          uploading: false
        };
      }
  handleUpload = async (options: any) => {
    const { file, onSuccess, onError } = options;
    this.setState({imageUrl: null, uploading: true})
    try {
      const base64Image = await this.readFileAsBase64(file);
      onSuccess('done', file);
      // console.log(base64Image);
      const url = 'http://0.0.0.0:8000/negative_image';
      const postData = {
        image: base64Image
      };

      axios
      .post(url, postData)
      .then((response) => {
        console.log(response.data.data)
        this.setState({imageUrl: response.data.data.image, uploading: false})
        message.success("Succesfully uploaded image")
      })
      .catch((error) => {
        message.error('Failed to upload image');
        console.error('Error:', error);
        this.setState({uploading: false})
        // Handle any error that occurred during the request
    });
    } catch (error) {
      console.error('Error:', error);
      onError(error);
      message.error('Failed to upload image');
      this.setState({uploading: false})
    }
  };

  readFileAsBase64 = (file: File): Promise<string> => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();

      reader.onload = () => {
        resolve(reader.result as string);
      };

      reader.onerror = (error) => {
        reject(error);
      };

      reader.readAsDataURL(file);
    });
  };

  render() {
    const {imageUrl, uploading} = this.state;
    return (
      <div className="my-component">
        <Upload
          name="image"
          // accept='</input type="image">'
          multiple={false}
          showUploadList={false}
          //action="/upload"
          customRequest={this.handleUpload}
          //beforeUpload={() => false} // Prevent immediate uploading
        >
          <Button icon={<UploadOutlined />} size="large">
            Click to Upload
          </Button>
        </Upload>
        {uploading && 
        <p style={{color: '#1890ff'}}>
            Uploading image...
        </p>}
        {imageUrl &&
        
          <div>
            <h4>Returned negative:</h4>
            <div className='upload-container'>
                <img
                src={imageUrl}
                alt="Selected"
                style={{
                    maxHeight: '90%',
                    maxWidth: '90%'
                }}
                />
            </div>
          </div>
          }
      </div>
    );
  }
}

export default NegativeImage;
