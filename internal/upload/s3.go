package upload

import (
  "context"
  "log"
  "os"
  "path/filepath"

  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
  "github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadDirectory(bucket_name string, dir_path string, region string) error {
  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
  if err != nil {
    log.Printf("error on loading default config\n%s\n", err)
    return err
  }

  s3_client := s3.NewFromConfig(cfg)
  uploader := manager.NewUploader(s3_client)

  err = filepath.WalkDir(dir_path, func(path string, info os.DirEntry, err error) error {
    if err != nil {
      return err
    }
    if info.IsDir() {
      return nil
    }

    key, err := filepath.Rel(dir_path, path)
    if err != nil {
      return err
    }

    file, err := os.Open(path)
    if err != nil {
      return err
    }
    defer file.Close()

    content_type := "text/html; charset=utf-8"
    _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
      Bucket: &bucket_name,
      Key: &key,
      Body: file,
      ContentType: &content_type,
    })
    if err != nil {
      log.Printf("error on uploading file %s\n%s\n", key, err)
      return err
    }
    return nil
  })
  if err != nil {
    log.Printf("error on walking directory %s\n%s\n", dir_path, err)
    return err
  }

  return nil
}
