// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.
 * 
 * For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).
 */
export function getObjectSignedUrl(args: GetObjectSignedUrlArgs): Promise<GetObjectSignedUrlResult> {
    return pulumi.runtime.invoke("gcp:storage/getObjectSignedUrl:getObjectSignedUrl", {
        "bucket": args.bucket,
        "contentMd5": args.contentMd5,
        "contentType": args.contentType,
        "credentials": args.credentials,
        "duration": args.duration,
        "extensionHeaders": args.extensionHeaders,
        "httpMethod": args.httpMethod,
        "path": args.path,
    });
}

/**
 * A collection of arguments for invoking getObjectSignedUrl.
 */
export interface GetObjectSignedUrlArgs {
    /**
     * The name of the bucket to read the object from
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64.
     * Typically retrieved from `google_storage_bucket_object.object.md5hash` attribute.
     * If you provide this in the datasource, the client (e.g. browser, curl) must provide the `Content-MD5` HTTP header with this same value in its request.
     */
    readonly contentMd5?: pulumi.Input<string>;
    /**
     * If you specify this in the datasource, the client must provide the `Content-Type` HTTP header with the same value in its request.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * What Google service account credentials json should be used to sign the URL. 
     * This data source checks the following locations for credentials, in order of preference: data source `credentials` attribute, provider `credentials` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable.
     */
    readonly credentials?: pulumi.Input<string>;
    /**
     * For how long shall the signed URL be valid (defaults to 1 hour - i.e. `1h`). 
     * See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.
     */
    readonly duration?: pulumi.Input<string>;
    /**
     * As needed. The server checks to make sure that the client provides matching values in requests using the signed URL. 
     * Any header starting with `x-goog-` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google.
     */
    readonly extensionHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * What HTTP Method will the signed URL allow (defaults to `GET`)
     */
    readonly httpMethod?: pulumi.Input<string>;
    /**
     * The full path to the object inside the bucket
     */
    readonly path: pulumi.Input<string>;
}

/**
 * A collection of values returned by getObjectSignedUrl.
 */
export interface GetObjectSignedUrlResult {
    /**
     * The signed URL that can be used to access the storage object without authentication.
     */
    readonly signedUrl: string;
}
