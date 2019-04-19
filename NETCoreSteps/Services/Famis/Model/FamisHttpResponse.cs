using System;
using System.Collections.Generic;
using System.Text;

namespace Famis.Model
{
    public class FamisHttpResponse
    {
        public bool Result { get; set; }
        public int Context { get; set; }
        public string Message { get; set; }
    }
}
