using System;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace Famis.Model
{
    public class ODataResponse<T>
    {
        public List<T> Value { get; set; }
        [JsonProperty("@odata.nextLink")]
        internal string NextLink { get; set; }

        public Uri NextLinkUrl {
            get {
                if (NextLink == null) {
                    return null;
                }
                var builder = new UriBuilder(NextLink) {Scheme = "https", Port = -1};
                return builder.Uri;
            }
        } 
    }
}