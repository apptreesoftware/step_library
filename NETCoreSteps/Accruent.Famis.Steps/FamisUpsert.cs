using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using StepCore;

namespace Accruent.Famis.Steps
{
    public abstract class FamisUpsert : ServiceStep
    {
        [Input(Description = "a json formatted object that you would like to create or update", Key = "Record")]
        public Dictionary<string, object> Object { get; set; }

        [Input(Description = "the endpoint of the entity you are trying to update")]
        public string Endpoint { get; set; }

        [Input(Description = "the id value of the record")]
        public int Id { get; set; }

        [Output(Description = "the success status of the upsert")]
        public bool Success { get; set; }

        [Output(Description = "the response message of the upsert")]
        public string Message { get; set; }

        [Output(Description = "the returned record from the upsert")]
        public object Record { get; set; }
    }
}