package adapters

//var awsConfig aws.Config
//var onceAwsConfig sync.Once

//var dynamodbClient *dynamodb.Client
//var onceDdbClient sync.Once

//var s3Client *s3.Client
//var onceS3Client sync.Once
//
//var sesClient *ses.Client
//var onceSesClient sync.Once
//
//var snsClient *sns.Client
//var onceSnsClient sync.Once

//var ssmClient *ssm.Client
//var onceSsmClient sync.Once

//func getAwsConfig() aws.Config {
//	cfg, err := awsConfigMod.LoadDefaultConfig(context.TODO())
//	if err != nil {
//		log.Fatalf("unable to load AWS SDK config, %v", err)
//	}
//	return cfg
//}
//
//func GetSsmClient() *ssm.Client {
//	onceSsmClient.Do(func() {
//		awsConfig = getAwsConfig()
//
//		region := config.Region
//
//		ssmClient = ssm.NewFromConfig(awsConfig, func(opt *ssm.Options) {
//			opt.Region = region
//		})
//	})
//
//	return ssmClient
//}

//func GetDynamodbClient() *dynamodb.Client {
//	onceDdbClient.Do(func() {
//		awsConfig = getAwsConfig()
//
//		region := config.Region
//
//		dynamodbClient = dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
//			opt.Region = region
//		})
//	})
//
//	return dynamodbClient
//}

//func GetS3Client() *s3.Client {
//	onceS3Client.Do(func() {
//		awsConfig = getAwsConfig()
//
//		region := config.Region
//
//		s3Client = s3.NewFromConfig(awsConfig, func(opt *s3.Options) {
//			opt.Region = region
//		})
//	})
//
//	return s3Client
//}

//func GetSesClient() *ses.Client {
//	onceSesClient.Do(func() {
//		awsConfig = getAwsConfig()
//
//		region := config.Region
//
//		sesClient = ses.NewFromConfig(awsConfig, func(opt *ses.Options) {
//			opt.Region = region
//		})
//	})
//
//	return sesClient
//}

//func GetSnsClient() *sns.Client {
//	onceSnsClient.Do(func() {
//		awsConfig = getAwsConfig()
//
//		region := config.Region
//
//		snsClient = sns.NewFromConfig(awsConfig, func(opt *sns.Options) {
//			opt.Region = region
//		})
//	})
//
//	return snsClient
//}
