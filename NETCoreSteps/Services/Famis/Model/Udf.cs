using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model {
    public class Udf {
        [JsonIgnore]
        public int? FieldId { get; set; }
        [JsonIgnore]
        public int? ListboxId { get; set; }
        [JsonIgnore]
        public string ListboxDescription { get; set; }
        public string Value { get; set; }
        public string FieldName { get; set; }
        [JsonIgnore]
        public string DataType { get; set; }
        [JsonIgnore]
        public string GroupDescription { get; set; }
        [JsonIgnore]
        public string ApplicationName { get; set; }
        [JsonIgnore]
        public int? TabOrder { get; set; }
        [JsonIgnore]
        public bool Required { get; set; }
        [JsonIgnore]
        public int? GroupTabOrder { get; set; }
    }
}
